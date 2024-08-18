import "./styles/ConversationPreview.css";

interface Conversation {
  chatName: string;
  timeOfLastMessage: Date;
  lastMessage: string;
}

interface ConversationPreviewProps {
  conversation: Conversation;
  onClick: () => void;
}

function ConversationPreview(props: ConversationPreviewProps) {
  return (
    <div className="conversation-container" onClick={props.onClick}>
      <div className="contacts-and-date-container">
        <h2>{props.conversation.chatName}</h2>
        <time>{props.conversation.timeOfLastMessage.toLocaleString()}</time>
      </div>
      <div className="message-container">{props.conversation.lastMessage}</div>
    </div>
  );
}

export default ConversationPreview;
