INSERT OR IGNORE INTO ollama_model_description (model_name, description)
VALUES
    ('llama3', 'Meta Llama 3: The most capable openly available LLM to date.'),
    ('gemma2', 'Gemma 2 is a family of lightweight open models from Google.');

INSERT OR IGNORE INTO ollama_model_tag (model_name, tag)
VALUES
    ('llama3', '8b'),
    ('llama3', '70b'),
    ('gemma2', '2b'),
    ('gemma2', '9b'),
    ('gemma2', '27b');
