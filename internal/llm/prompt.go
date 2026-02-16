package llm

func BuildPrompt(requirement string) string {
	return `
		You are a senior Backend engineer.

		Generate a production-ready project for the following requirement:

		` + requirement + `

		Rules:
		- Keep it as per industry standards.
		- Use comments to explain complex logic.
		- Include error handling and logging.
		- Include tests for critical components.
		Return ONLY valid JSON in this format:

		{
		"files": [
			{
			"path": "file_path/name here",
			"content": "file content here"
			}
		]
		}
		`
}