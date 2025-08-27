FROM ollama/ollama:latest

ENV OLLAMA_HOME=/root/.ollama

VOLUME ["/root/.ollama"]

EXPOSE 11434

ENTRYPOINT ["/bin/sh", "-c"]

CMD ["ollama serve & sleep 5 && ollama pull gemma3:1b && wait"]
