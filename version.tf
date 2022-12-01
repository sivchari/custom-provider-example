terraform {
  required_providers {
    example = {
      version = "~> 1.0.0"
      source  = "github.com/sivchari/example"
    }
  }
}

provider example {}
