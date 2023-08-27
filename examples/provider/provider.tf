terraform {
  required_providers {
    MyTest = {
      source  = "testing/MyTest"
      version = "0.3.5"
    }
  }
}

provider "MyTest" {
  # Configuration options
}