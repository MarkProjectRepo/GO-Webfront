from transformers import pipeline

class Pipeline():

	def __init__(self):
		self.pipe = pipeline("zero-shot-classification")


	def predict(self, payload):
		classes = payload["classes"].split(",")
		return self.pipe(payload["text"], classes)