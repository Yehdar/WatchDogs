import React, { useState } from "react";
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
    <div
      className="App"
      class="flex items-center justify-center min-h-screen bg-slate-600"
    >
      <div className="text-center">
        <h1 class="text-3xl text-white font-bold mb-4">
          Video Streaming Platform
        </h1>
        <video width="700px" height="400px" controls>
          <source
            src="https://d2i8zdo0i9xncc.cloudfront.net/Kobe%20Bryant%20&%20Kanye%20West%20.mp4"
            type="video/mp4"
          />
        </video>
        <div className="mt-4">
          <h2 className="text-xl text-slate-400 font-bold mb-2">Comments</h2>
          <u1 className="list-disc text-slate-100">
            {comments.map((comment, index) => (
              <li key={index}>{comment}</li>
            ))}
          </u1>
          <div className="mt-4">
            <textarea
              className="w-full p-2 border rounded bg-slate-100"
              placeholder="Add a comment..."
              value={newComment}
              onChange={(e) => setNewComment(e.target.value)}
            ></textarea>
            <button
              className="mt-2 bg-cyan-500 text-white p-2 rounded"
              onClick={addComment}
            >
              Add Comment
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
