import "./styles/MessagesView.css";
import { faker } from "@faker-js/faker";

interface Message {
  isSentByCurrentUser: boolean;
  timestamp: Date;
  text: string;
}

function CreateDummyMessage() {
  const isSentByCurrentUser = Math.random() > 0.5;
  const timestamp = faker.date.recent();
  const text = faker.lorem.sentences();

  return {
    isSentByCurrentUser: isSentByCurrentUser,
    timestamp: timestamp,
    text: text,
  };
}

function MessageView() {
  const messages: Message[] = [];
  for (let i = 0; i < 20; i++) {
    messages.push(CreateDummyMessage());
  }

  messages.sort((a, b) => a.timestamp.getTime() - b.timestamp.getTime());
  return (
    <div className="messages-container">
      {messages.map((message) => (
        <div className={message.isSentByCurrentUser ? "sent" : "received"}>
          <div style={{ padding: "20px" }}>{message.text}</div>
        </div>
      ))}
    </div>
  );
}

export default MessageView;
