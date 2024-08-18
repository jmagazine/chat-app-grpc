import { useState, useEffect, useRef } from "react";
import "./styles/MessagesView.css";
import { faker } from "@faker-js/faker";
import { TextField } from "@mui/material";

interface Message {
  isSentByCurrentUser: boolean;
  timestampInMS: number;
  text: string;
}

function CreateDummyMessage(): Message {
  const isSentByCurrentUser = Math.random() > 0.5;
  const timestamp = faker.date.recent().getTime();
  const text = faker.lorem.sentences();

  return {
    isSentByCurrentUser: isSentByCurrentUser,
    timestampInMS: timestamp,
    text: text,
  };
}

function MessagesView() {
  const [messages, setMessages] = useState<Message[]>([]);

  useEffect(() => {
    const initialMessages: Message[] = [];
    for (let i = 0; i < 20; i++) {
      initialMessages.push(CreateDummyMessage());
    }
    initialMessages.sort((a, b) => a.timestampInMS - b.timestampInMS);
    setMessages(initialMessages);
  }, []);

  const handleSendMessage = () => {
    const text = (
      document.getElementById("filled-multiline-static") as HTMLInputElement
    )?.value;
    if (!text || text.length === 0) {
      return;
    }
    const messageToSend: Message = {
      isSentByCurrentUser: true,
      timestampInMS: Date.now(),
      text: text,
    };

    setMessages((prevMessages) => [messageToSend, ...prevMessages]);

    // Clear the text field after sending
    (
      document.getElementById("filled-multiline-static") as HTMLInputElement
    ).value = "";

    // Scroll to the bottom after sending a message
  };

  useEffect(() => {
    const handleKeyPress = (e: KeyboardEvent) => {
      if (e.key === "Enter") {
        handleSendMessage();
      }
    };

    document.addEventListener("keypress", handleKeyPress);

    // Cleanup event listener on component unmount
    return () => {
      document.removeEventListener("keypress", handleKeyPress);
    };
  }, []);

  return (
    <div className="messages-container">
      {messages.map((message) => (
        <div
          key={message.timestampInMS}
          className={message.isSentByCurrentUser ? "sent" : "received"}
        >
          <div style={{ padding: "20px" }}>{message.text}</div>
        </div>
      ))}
      <div className="textfield-container">
        <TextField
          id="filled-multiline-static"
          label="Multiline"
          multiline
          rows={4}
          placeholder="Start typing..."
          variant="filled"
          sx={{ width: "65vw" }}
        />
      </div>
    </div>
  );
}

export default MessagesView;
