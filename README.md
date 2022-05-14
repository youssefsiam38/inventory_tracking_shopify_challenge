# Inventory Tracking - Shopify Intern Challenge 2022

<img align="right" width="159px" src="/assets/shopify.jpg">

This project built from scratch using the [go](https://go.dev/) programming language without taking advantage of any [batteries-included framework](https://en.wiktionary.org/wiki/batteries-included), with the try to apply the [uncle bob's clean architecture principles](https://www.amazon.com/Clean-Architecture-Craftsmans-Software-Structure/dp/0134494164).

## Contents

- [How to use the project](#how-to-use-the-project)
  - [Installation](#installation)
  - [Features included](#features-included)
  - [Technical aspects](#technical-aspects)
    - [Two database repositories (inmemory/replit)](#two-database-repositories)
    - [Why go1.14?](#why-go114)


## Installation

To install The project and see the code, you need to install Go and set your Go workspace first.

1. First you need to clone this repository in you machine

2. Then you will need [Go](https://golang.org/) installed (**version 1.14+ is required**), then you can use the below Go command to run the project.

```sh
$ go run main.go -db inmemory
```

3. Go to http://localhost:8080 to open the project in the browser

## Features Included

### Steps to use each featrue
- View a list of items
    1. Open the [home page](https://shopifyintern.youssefsiam.me/inventory/)
- Create inventory items
    1. Click on Create item
    1. Fill the form
- Edit item
    1. Click on Edit icon in the row you want to change
    1. Fill the form
- Delete item and allow deletion comments
    1. Click on Delete icon in the row you want to change
    1. Write the comment of the deletion
- View a list of deleted items
    1. Click on Deleted items
- Undeletion
    1. Click on Deleted items
    1. Click on Edit icon in the row you want to change
    1. click on Undelete(you are not allowed to edit deleted item unless you undelete it)


## Technical Aspects

### Two Database Repositories

This project's architecture use two repositories that implements the same interface, you just need to specify the `-db` flag (replit is the default):-

##### Replit Repository 

To take advantage of the [replit database](https://docs.replit.com/hosting/database-faq) in production to avoid exposing the database queries to external database to optimize the performance.

##### Inmermory Repository 

To use it for development and in the test cases.

### Why Go1.14?

Because this the version used in the replit's Go workspace.