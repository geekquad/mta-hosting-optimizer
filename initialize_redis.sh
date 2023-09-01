#!/bin/sh

# Set data in Redis using HSET commands
redis-cli HSET mta-prod-1 127.0.0.1 true 127.0.0.2 false
redis-cli HSET mta-prod-2 127.0.0.3 true 127.0.0.4 true 127.0.0.5 false
redis-cli HSET mta-prod-3 127.0.0.6 false