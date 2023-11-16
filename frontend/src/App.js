import React, { useState } from "react";
import logo from "./logo.svg";
import "./App.css";

function App() {
  const [comments, setComments] = useState([]);
  const [newComment, setNewComment] = useState("");

  const addComment = () => {
    if (newComment.trim() !== "") {
      setComments([...comments, newComment]);
      setNewComment("");
    }
  };

  return (
    <div className="App" class="flex items-center justify-center min-h-screen">
      <div className="text-center">
        <h1 class="text-3xl font-bold mb-4">Video Streaming Platform</h1>
        <video width="700px" height="400px" controls>
          <source
            src="https://d2i8zdo0i9xncc.cloudfront.net/Kobe%20Bryant%20&%20Kanye%20West%20.mp4"
            type="video/mp4"
          />
        </video>
        <div>
          <h2>Comments</h2>
          <u1>
            {comments.map((comment, index) => (
              <li key={index}>{comment}</li>
            ))}
          </u1>
          <div>
            <textarea placeholder="aaaaa">
              <button>Add Comment</button>
            </textarea>
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
