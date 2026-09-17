from openai import OpenAI

# GEMINI WITH OPENAI ROUTER

client = OpenAI(
    api_key="AIzaSyCx51aWji7MgTY6_hGwIm9IjJkU_5EJnzM",
    base_url="https://generativelanguage.googleapis.com/v1beta/openai/",
)

response = client.chat.completions.create(
    model="gemini-2.5-flash",
    messages=[
        {"role": "user", "content": "Hi, How are you"},
    ],
)

print(response.choices[0].message.content)
