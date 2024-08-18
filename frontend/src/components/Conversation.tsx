import React from "react";
import "./styles/Conversation.css";

interface Conversation {
  contact: string;
  timeOfLastMessage: Date;
  lastMessage: string;
}

function ConversationComponent({
  conversation,
}: {
  conversation: Conversation;
}) {
  return (
    <div className="conversation-container">
      <div className="contacts-and-date-container">
        <h2>{conversation.contact}</h2>
        <time>{conversation.timeOfLastMessage.toLocaleString()}</time>
      </div>
      <div className="message-container">{conversation.lastMessage}</div>
    </div>
  );
}

export default ConversationComponent;
