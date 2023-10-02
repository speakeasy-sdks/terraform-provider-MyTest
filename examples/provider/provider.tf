terraform {
  required_providers {
    MyTest = {
      source  = "testing/MyTest"
      version = "0.9.1"
    }
  }
}

provider "MyTest" {
  # Configuration options
}