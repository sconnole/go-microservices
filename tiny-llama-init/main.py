import time
import requests
import os

OLLAMA_URL = os.getenv("OLLAMA_URL", "http://ollama-service.default.svc.cluster.local:11434")

# Wait for Ollama to be ready
print("Waiting for Ollama to become available...")
while True:
    try:
        response = requests.get(OLLAMA_URL)
        if response.status_code == 200:
            print("Ollama is ready.")
            break
    except requests.exceptions.ConnectionError:
        pass
    print("checking again")
    time.sleep(1)

data = '{"name": "tinyllama"}'  # raw JSON string

headers = {
    'Content-Type': 'application/json',
}

response = requests.post(f"{OLLAMA_URL}/api/pull", data=data, headers=headers)

print(response.text)

