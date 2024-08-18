// import { ChatMessage } from "../gen/es/chat_pb";

// interface ChatHistory {
//   chatName: string;
//   chatIcon: HTMLImageElement;
//   chatMessages: ChatMessage[];
// }

import ConversationComponent from "./Conversation";
import MessageView from "./MessagesView";
import "./styles/ConversationsView.css";
import { faker } from "@faker-js/faker";
// lazily eval inbox

function CreateDummyConversation() {
  const contact = faker.person.fullName();
  const timeOfLastMessage = faker.date.recent({ days: 5 });
  const lastMessage = faker.lorem.sentence();

  return {
    contact: contact,
    timeOfLastMessage: timeOfLastMessage,
    lastMessage: lastMessage,
  };
}

function ConversationsView() {
  const conversationComponents = [];

  for (let i = 0; i < 20; i++) {
    conversationComponents.push(
      <ConversationComponent
        conversation={CreateDummyConversation()}
      ></ConversationComponent>
    );
  }
  conversationComponents.sort(
    (a, b) =>
      b.props.conversation.timeOfLastMessage.getTime() -
      a.props.conversation.timeOfLastMessage.getTime()
  );

  return (
    <div className="conversations-view-main-container">
      <div className="contacts-container">
        <h2>Contacts</h2>

        {conversationComponents}
      </div>
      <div className="conversations-container">
        <h2>Messages</h2>
        <MessageView />
      </div>
    </div>
  );
}

export { ConversationComponent };
export default ConversationsView;
