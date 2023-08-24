terraform {
  required_providers {
    MyTest = {
      source  = "testing/MyTest"
      version = "0.2.2"
    }
  }
}

provider "MyTest" {
  # Configuration options
}