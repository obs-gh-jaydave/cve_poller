# CVE Poller

This project contains the necessary files to set up an HTTP poller using Observe to connect to an open-source CVE database.

## Folder Structure

cve_poller/
├── LICENSE
├── Makefile
├── Overview.md
├── README.md
├── config
│   └── cve_poller.yaml
├── cve_data
│   ├── allitems.json
│   └── allitems.xml
├── example.tf
├── go.mod
├── go.sum
├── main.go
├── main.tf
├── mainifest.yaml
├── outputs.tf
├── sources
│   ├── external
│   │   ├── README.md
│   │   └── metadata.yaml
│   └── poller
│       ├── README.md
│       ├── main.tf
│       ├── metadata.yaml
│       ├── outputs.tf
│       ├── variables.tf
│       └── versions.tf
└── tftests
    ├── README.md
    └── default
        ├── main.tf
        ├── outputs.tf
        ├── variables.tf
        └── versions.tf

8 directories, 27 files

## Setup

1. Clone the repository:
    ```
    git clone <repository_url>
    cd cve_poller
    ```

2. Build and install the crawler:
    ```
    go build -o poll_crawler
    ```

3. Run the poller:
    ```
    ./poll_crawler
    ```

## Running the Poller

To run the poller locally, execute the following script:
```
./poll_crawler
```

## Terraform

To set up the poller using Terraform, navigate to the `cve_poller` directory and apply the configuration:

1. Initialize Terraform:
   ```
   terraform init
   ```
2. Run Terraform Plan w/ debugging enabled:
   ```
   TF_LOG=DEBUG terraform plan
   ```
3. Apply the Terraform configuration:
   ```
   terraform apply
   ```
## Notes
- The CVE poller fetches data from an open-source CVE database and saves it in both `XML` and `JSON` formats in the `cve_data directory`.
- Ensure that you have Go installed and properly configured on your machine to build and run the poller.
- Terraform must be installed and configured to manage the infrastructure as code.
