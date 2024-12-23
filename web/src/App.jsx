import { useState } from 'react'

function App() {
  const [input, setInput] = useState('')
  const [result, setResult] = useState('')
  const [error, setError] = useState('')

  const handleSubmit = async () => {
    try {
      const response = await fetch('/api/minify', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ input }),
      })
      
      const data = await response.json()
      
      if (!response.ok) {
        setError(data.error)
        setResult('')
      } else {
        setResult(data.result)
        setError('')
      }
    } catch (err) {
      setError('Failed to process request')
      setResult('')
    }
  }

  return (
    <div className="container mx-auto p-4 max-w-2xl">
      <h1 className="text-2xl font-bold mb-4">JSON Minifier</h1>
      <div className="space-y-4">
        <div>
          <label className="block mb-2">Input JSON:</label>
          <textarea
            className="w-full h-48 p-2 border rounded"
            value={input}
            onChange={(e) => setInput(e.target.value)}
            placeholder="Paste your JSON here..."
          />
        </div>
        
        <button
          onClick={handleSubmit}
          className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
        >
          Minify
        </button>

        {error && (
          <div className="text-red-500 mt-4">
            {error}
          </div>
        )}

        {result && (
          <div>
            <label className="block mb-2">Result:</label>
            <pre className="w-full p-2 bg-gray-100 rounded overflow-x-auto">
              {result}
            </pre>
          </div>
        )}
      </div>
    </div>
  )
}

export default App