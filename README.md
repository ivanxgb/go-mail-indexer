# Go Mail-Indexer

## Description

***Go Mail-Indexer*** is an application that given a directory route of mail files, it will search through the files and 
index them into a search engine database (Zinc Search Engine).

Then, from the frontend, the user can search through the indexed files using a backend API also written in Go. 
The API will return the results of Mails that match the search query.

## Project Details

### Tech Stack

- Go 1.20 - [Go](https://go.dev/): Used for the backend API and the indexer
- Vue 3 - [Vue](https://v3.vuejs.org/): Frontend framework
- Tailwind CSS - [Tailwind](https://tailwindcss.com/): CSS framework
- Search Engine Database - [Zinc Search Engine](https://github.com/zincsearch/zincsearch/): Search engine database used for indexing and searching

### Resources

- Mail Database: [Enron Mail Database](http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz)
