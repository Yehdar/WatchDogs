import logo from "./logo.svg";
import "./App.css";

function App() {
  return (
    <div className="App">
      <h1>Video Streaming Platform</h1>
      <video width="700px" height="400px" controls>
        <source
          src="https://d2i8zdo0i9xncc.cloudfront.net/Kobe%20Bryant%20&%20Kanye%20West%20.mp4"
          type="video/mp4"
        />
      </video>
    </div>
  );
}

export default App;
