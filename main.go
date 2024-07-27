package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v2"
)

type CVEData struct {
	Items []CVEItem `xml:"item"`
}

type CVEItem struct {
	ID               string            `xml:"name,attr" json:"ID"`
	Description      string            `xml:"desc" json:"Description"`
	References       []Reference       `xml:"refs>ref" json:"References"`
	Status           string            `xml:"status" json:"Status"`
	Phase            Phase             `xml:"phase" json:"Phase"`
	Votes            []Vote            `xml:"votes>modify" json:"Votes"`
	Comments         []Comment         `xml:"comments>comment" json:"Comments"`
	CWE              string            `xml:"cwe" json:"CWE"`
	CVSS             CVSS              `xml:"cvss" json:"CVSS"`
	Configurations   []Configuration   `xml:"configurations>configuration" json:"Configurations"`
	Impact           Impact            `xml:"impact" json:"Impact"`
	Advisories       []Advisory        `xml:"advisories>advisory" json:"Advisories"`
	Workarounds      []Workaround      `xml:"workarounds>workaround" json:"Workarounds"`
	VendorStatements []VendorStatement `xml:"vendorStatements>vendorStatement" json:"VendorStatements"`
}

type Reference struct {
	Source string `xml:"source,attr" json:"Source"`
	URL    string `xml:"url,attr" json:"URL"`
}

type Phase struct {
	Date string `xml:"date,attr" json:"Date"`
	Text string `xml:",chardata" json:"Text"`
}

type Vote struct {
	Count int    `xml:"count,attr" json:"Count"`
	Text  string `xml:",chardata" json:"Text"`
}

type Comment struct {
	Voter string `xml:"voter,attr" json:"Voter"`
	Text  string `xml:",chardata" json:"Text"`
}

type CVSS struct {
	BaseScore          float64 `xml:"base_score" json:"BaseScore"`
	Exploitability     float64 `xml:"exploitability" json:"Exploitability"`
	Impact             float64 `xml:"impact" json:"Impact"`
	TemporalScore      float64 `xml:"temporal_score" json:"TemporalScore"`
	EnvironmentalScore float64 `xml:"environmental_score" json:"EnvironmentalScore"`
}

type Configuration struct {
	ID       string `xml:"id,attr" json:"ID"`
	Operator string `xml:"operator,attr" json:"Operator"`
}

type Impact struct {
	Confidentiality string `xml:"confidentiality" json:"Confidentiality"`
	Integrity       string `xml:"integrity" json:"Integrity"`
	Availability    string `xml:"availability" json:"Availability"`
}

type Advisory struct {
	Source string `xml:"source,attr" json:"Source"`
	URL    string `xml:"url,attr" json:"URL"`
}

type Workaround struct {
	Text string `xml:",chardata" json:"Text"`
}

type VendorStatement struct {
	Vendor string `xml:"vendor,attr" json:"Vendor"`
	Text   string `xml:",chardata" json:"Text"`
}

func main() {
	log.Println("Starting CVE poller")

	// Load configuration from YAML file
	config, err := loadConfig("config/cve_poller.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Create a directory to store CVE data files if it doesn't exist
	cveDataDir := "cve_data"
	if _, err := os.Stat(cveDataDir); os.IsNotExist(err) {
		err = os.Mkdir(cveDataDir, 0755)
		if err != nil {
			log.Fatalf("Failed to create directory: %v", err)
		}
	}

	// Download the CVE data file
	resp, err := http.Get(config.Requests[0].URL)
	if err != nil {
		log.Fatalf("Failed to download CVE data: %v", err)
	}
	defer resp.Body.Close()

	// Read the downloaded file
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read CVE data: %v", err)
	}

	// Save the downloaded file locally
	cveFilePath := filepath.Join(cveDataDir, "allitems.xml")
	err = ioutil.WriteFile(cveFilePath, data, 0644)
	if err != nil {
		log.Fatalf("Failed to save CVE data: %v", err)
	}
	log.Printf("CVE data saved to %s\n", cveFilePath)

	// Process the CVE data
	var cveData CVEData
	reader := strings.NewReader(string(data))
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&cveData)
	if err != nil {
		log.Fatalf("Failed to parse CVE data: %v", err)
	}

	// Convert the CVE data to JSON format
	jsonData, err := json.MarshalIndent(cveData.Items, "", "  ")
	if err != nil {
		log.Fatalf("Failed to convert CVE data to JSON: %v", err)
	}

	// Save the JSON data locally
	jsonFilePath := filepath.Join(cveDataDir, "allitems.json")
	err = ioutil.WriteFile(jsonFilePath, jsonData, 0644)
	if err != nil {
		log.Fatalf("Failed to save CVE JSON data: %v", err)
	}
	log.Printf("CVE JSON data saved to %s\n", jsonFilePath)

	log.Println("CVE poller completed successfully")
}

type Config struct {
	Requests []struct {
		URL    string            `yaml:"url"`
		Params map[string]string `yaml:"params"`
	} `yaml:"requests"`
}

func loadConfig(configPath string) (*Config, error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
