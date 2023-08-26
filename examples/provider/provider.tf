terraform {
  required_providers {
    MyTest = {
      source  = "testing/MyTest"
      version = "0.3.2"
    }
  }
}

provider "MyTest" {
  # Configuration options
}