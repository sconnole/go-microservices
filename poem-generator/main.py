import requests 
import os

RANDOM_JOB_URL = os.getenv("RANDOM_JOB_URL", "http://api.default.svc.cluster.local/random-job")
OLLAMA_URL = os.getenv("OLLAMA_URL", "http://ollama-service.default.svc.cluster.local:11434")
OLLAMA_MODEL = os.getenv("OLLAMA_MODEL", "tinyllama")

def main(): 
    try:
        res = requests.get(RANDOM_JOB_URL)
        res.raise_for_status()
        words = res.text.strip()
        if not words:
            print("Received empty words from random-job.")
            return
    except Exception as e:
        print(f"Failed to get words from {RANDOM_JOB_URL}: {e}")
        return

    print(f"\nTopic: {words}")
    
    prompt = f"Write a short poem inspired by these works: \"{words}\""

    try:
        res = requests.post(
            f"{OLLAMA_URL}/api/generate",
            json={
                "model": OLLAMA_MODEL,
                "prompt": prompt,
                "stream": False
            }
        )
        res.raise_for_status()
        poem = res.json().get("response", "").strip()
        print(f"\nPoem:\n{poem}")
    except Exception as e:
        print(f"Error generating poem: {e}")

if __name__ == "__main__":
    main()