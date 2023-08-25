terraform {
  required_providers {
    MyTest = {
      source  = "testing/MyTest"
      version = "0.2.3"
    }
  }
}

provider "MyTest" {
  # Configuration options
}