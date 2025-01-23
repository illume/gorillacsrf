import { useState } from "react";
import reactLogo from "./assets/react.svg";
import "./App.css";



// Fetch the CSRF token from the headers
async function getCsrfToken() {
	const response = await fetch('/api/user/1');
	const csrfToken = response.headers.get('X-CSRF-Token');
	return csrfToken;
}

async function makeRequest(endpoint: string, formData: Record<string, unknown>) {
	try {
		const csrfToken = await getCsrfToken();
    alert(csrfToken);
    const headers = {
      'Content-Type': 'application/json',
      'X-CSRF-Token': csrfToken || ''
    };

		const response = await fetch(`/${endpoint}`, {
			method: 'POST',
			headers: headers,
			body: JSON.stringify(formData),
      credentials: 'include',
		});

    alert(response.status + response.statusText);
    if (!response.ok) {
      console.error('Fetch error:', response.statusText);
      throw new Error('Network response was not ok');
    }
		const result = await response.json();
    alert(JSON.stringify(result));

	} catch (err) {
		// Handle the exception
		console.error('Error:', err);
	}
}

const endpoint = 'api/number';
const formData = { key: 'value' };


function App() {
	const [count, setCount] = useState(0);

	return (
		<div className="App">
			{/* <form>
				<input type="hidden" name="csrftokenxxxx" value="hidden" />
				<input type="text" placeholder="Type something" />
				<button type="submit">Submit</button>
			</form> */}
      <button onClick={() => makeRequest(endpoint, formData)}>Make Request</button>

			<div>
				<a href="https://reactjs.org" target="_blank" rel="noreferrer">
					<img src={reactLogo} className="logo react" alt="React logo" />
				</a>
			</div>
			<h1>Rspack + React + TypeScript</h1>
			<div className="card">
				<button type="button" onClick={() => setCount(count => count + 1)}>
					count is {count}
				</button>
				<p>
					Edit <code>src/App.tsx</code> and save to test HMR
				</p>
			</div>
			<p className="read-the-docs">
				Click on the Rspack and React logos to learn more
			</p>
		</div>
	);
}

export default App;
