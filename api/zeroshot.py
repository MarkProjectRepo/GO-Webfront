from transformers import pipeline

class PythonPredictor:
	def __init__(self, config):
            self.pipe = pipeline('sentiment-analysis')

	def predict(self, payload):
            return self.pipe(payload["text"], min_length=5, max_length=200)
