terraform {
  required_providers {
    MyTest = {
      source  = "testing/MyTest"
      version = "0.2.0"
    }
  }
}

provider "MyTest" {
  # Configuration options
}