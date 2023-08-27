terraform {
  required_providers {
    MyTest = {
      source  = "testing/MyTest"
      version = "0.3.4"
    }
  }
}

provider "MyTest" {
  # Configuration options
}