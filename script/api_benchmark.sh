#!/bin/bash
# API Benchmark Test Script - Fixed parameter issues

# Basic parameters
BASE_URL="http://localhost:8080"  # Replace with your actual service address
REQUESTS=10000                    # Total number of requests
CONCURRENCY=100                   # Concurrency level
RESULTS_DIR="./"                  # Results directory (current directory)

# Create results directory
mkdir -p $RESULTS_DIR
TIMESTAMP=$(date +"%Y%m%d_%H%M%S")
RESULT_FILE="$RESULTS_DIR/benchmark_$TIMESTAMP.txt"
LATEST_FILE="$RESULTS_DIR/latest.txt"

# Use a valid Ethereum address
# Note: This is a sample address, make sure it has relevant records in your system
VALID_ADDRESS="0x71C7656EC7ab88b098defB751B7401B5f6d8976F"

# Ensure all required query parameters exist
PAGE="1"
PAGE_SIZE="10"
ORDER="desc"
CHAIN_ID="11155420"  # Use a valid chain ID

# Record test start
echo "=== API Performance Test Started - $(date) ===" | tee -a $RESULT_FILE

# Test health check endpoint
echo -e "\n\n=== Testing Health Check Endpoint ===" | tee -a $RESULT_FILE
hey -n $REQUESTS -c $CONCURRENCY $BASE_URL/healthz | tee -a $RESULT_FILE

# Test staking records endpoint
echo -e "\n\n=== Testing Staking Records Endpoint ===" | tee -a $RESULT_FILE
hey -n $REQUESTS -c $CONCURRENCY "$BASE_URL/api/v1/staking-records?address=$VALID_ADDRESS&page=$PAGE&pageSize=$PAGE_SIZE&order=$ORDER" | tee -a $RESULT_FILE

# Test bridge records endpoint
echo -e "\n\n=== Testing Bridge Records Endpoint ===" | tee -a $RESULT_FILE
hey -n $REQUESTS -c $CONCURRENCY "$BASE_URL/api/v1/bridge-records?address=$VALID_ADDRESS&page=$PAGE&pageSize=$PAGE_SIZE&order=$ORDER" | tee -a $RESULT_FILE

# Test staking validity check endpoint
echo -e "\n\n=== Testing Staking Validity Check Endpoint ===" | tee -a $RESULT_FILE
hey -n $REQUESTS -c $CONCURRENCY "$BASE_URL/api/v1/staking-valid?address=$VALID_ADDRESS" | tee -a $RESULT_FILE

# Test bridge validity check endpoint
echo -e "\n\n=== Testing Bridge Validity Check Endpoint ===" | tee -a $RESULT_FILE
hey -n $REQUESTS -c $CONCURRENCY "$BASE_URL/api/v1/bridge-valid?address=$VALID_ADDRESS" | tee -a $RESULT_FILE

# Test bridge gas fee endpoint
echo -e "\n\n=== Testing Bridge Gas Fee Endpoint ===" | tee -a $RESULT_FILE
hey -n $REQUESTS -c $CONCURRENCY "$BASE_URL/api/v1/bridge-gas-fee?chain_id=$CHAIN_ID" | tee -a $RESULT_FILE

# Record test end
echo -e "\n\n=== API Performance Test Completed - $(date) ===" | tee -a $RESULT_FILE

# Update latest result link
cp $RESULT_FILE $LATEST_FILE

echo "Test completed! Results saved to: $RESULT_FILE"