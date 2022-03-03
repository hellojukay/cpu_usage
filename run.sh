#!/bin/bash

cd frontend
yarn build
cd ..
wails dev
