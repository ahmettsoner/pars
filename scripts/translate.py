import openai
import os
import sys

def translate_text(text, target_language):
    response = openai.ChatCompletion.create(
        model="gpt-4",
        messages=[
            {"role": "system", "content": "You are a helpful assistant."},
            {"role": "user", "content": f"Translate the following text to {target_language}:\n\n{text}"}
        ]
    )
    return response.choices[0].message['content'].strip()

def read_file(file_path):
    with open(file_path, 'r', encoding='utf-8') as f:
        return f.read()

def write_file(file_path, content):
    with open(file_path, 'w', encoding='utf-8') as f:
        f.write(content)

if __name__ == "__main__":
    if len(sys.argv) < 4:
        print("Usage: python translate.py <source_path> <target_path> <target_language>")
        sys.exit(1)

    source_path = sys.argv[1]
    target_path = sys.argv[2]
    target_language = sys.argv[3]

    openai.api_key = os.getenv("OPENAI_API_KEY")

    text_to_translate = read_file(source_path)
    translated_text = translate_text(text_to_translate, target_language)

    write_file(target_path, translated_text)