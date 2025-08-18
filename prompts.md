Question: %s
Answer options:
%s

Check the question, answer options, and explanations for any errors.
⚠️ Explanations are allowed only for correct answers (isCorrect: true).

- If there are errors in the question, answers, or explanations, correct them and return the fixed JSON.
- If everything is correct, return an empty JSON: {}.

⚠️ The response must be **JSON only**:
{
  "question": "corrected question",
  "options": [
    {"text": "answer text", "isCorrect": true/false, "explanation": "explanation"}
  ]
}

Do not include any extra text, comments, or reasoning. Return only the JSON object.
