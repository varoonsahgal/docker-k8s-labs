"""Fortune Cookie API: dispenses container wisdom, one cookie at a time."""
import os
import random
import socket

from flask import Flask, jsonify

app = Flask(__name__)

FORTUNES = [
    "A container that works on your machine will one day meet production.",
    "You will find great joy in a 5 MB image.",
    "The tag 'latest' is a lie you tell yourself.",
    "An unpinned base image brings surprises at 2am.",
    "He who stores data in a container without a volume, stores it in the wind.",
    "Localhost inside a container is not the localhost you seek.",
    "Your next deploy will be rolled out gradually and without fear.",
    "A healthcheck today keeps the pager away.",
    "Kubernetes is in your future. Bring snacks.",
    "Many replicas make light work.",
]


@app.get("/")
def fortune():
    return jsonify(
        fortune=random.choice(FORTUNES),
        served_by=socket.gethostname(),
        lucky_port=random.randint(1024, 65535),
    )


@app.get("/healthz")
def healthz():
    return "ok\n"


if __name__ == "__main__":
    # 🕵️ Something about this line is going to hurt you inside a container...
    app.run(host="127.0.0.1", port=int(os.environ.get("PORT", 5000)))
