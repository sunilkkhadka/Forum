import { useEffect } from "react";
import "./App.css";

function App() {
  useEffect(() => {
    const healthCheck = async () => {
      const response = await fetch("http://localhost:8080/api/v1/healthcheck");

      const data = await response.json();

      console.log(data);
    };

    healthCheck();
  });

  return (
    <main>
      <h1>Hello Blog</h1>
    </main>
  );
}

export default App;
