import { useState } from "react";

const defaultMessage = "Click the button to load the backend greeting.";

export default function App() {
  const [message, setMessage] = useState(defaultMessage);
  const [loading, setLoading] = useState(false);

  const loadGreeting = async () => {
    setLoading(true);
    try {
      const response = await fetch("http://localhost:8080/api/health");
      const data = await response.json();
      setMessage(data.message ?? defaultMessage);
    } catch (error) {
      setMessage("We could not reach the backend yet.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="page">
      <header className="hero">
        <p className="eyebrow">Security Cars</p>
        <h1>Simple Web Starter</h1>
        <p className="subtitle">
          This is a minimal React interface that talks to the Go backend.
        </p>
        <button className="primary-button" onClick={loadGreeting} disabled={loading}>
          {loading ? "Loading..." : "Get backend greeting"}
        </button>
      </header>
      <section className="card">
        <h2>Backend response</h2>
        <p>{message}</p>
      </section>
      <section className="card">
        <h2>Next steps</h2>
        <ul>
          <li>Run the Go API to keep the message dynamic.</li>
          <li>Edit App.jsx to add your own sections.</li>
          <li>Style changes are in styles.css.</li>
        </ul>
      </section>
    </div>
  );
}
