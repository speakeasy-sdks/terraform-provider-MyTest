terraform {
  required_providers {
    MyTest = {
      source  = "testing/MyTest"
      version = "0.10.1"
    }
  }
}

provider "MyTest" {
  # Configuration options
}