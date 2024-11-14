INSERT OR IGNORE INTO model_description (name, description)
VALUES ('llama3', 'Meta Llama 3: The most capable openly available LLM to date.');

INSERT OR IGNORE INTO model_description (name, description)
VALUES ('gemma2', 'Gemma 2 is a family of lightweight open models from Google.');

INSERT OR IGNORE INTO model_tag (model, tag)
VALUES ('llama3', '8b');

INSERT OR IGNORE INTO model_tag (model, tag)
VALUES ('llama3', '70b');

INSERT OR IGNORE INTO model_tag (model, tag)
VALUES ('gemma2', '2b');

INSERT OR IGNORE INTO model_tag (model, tag)
VALUES ('gemma2', '9b');

INSERT OR IGNORE INTO model_tag (model, tag)
VALUES ('gemma2', '27b');
