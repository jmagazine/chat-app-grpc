// import { ChatMessage } from "../gen/es/chat_pb";

// interface ChatHistory {
//   chatName: string;
//   chatIcon: HTMLImageElement;
//   chatMessages: ChatMessage[];
// }

import { useEffect, useMemo, useState } from "react";
import ConversationPreview from "./ConversationPreview";
import MessagesView from "./MessagesView";
import "./styles/HomeView.css";
import { faker } from "@faker-js/faker";
// lazily eval inbox

function CreateDummyConversation() {
  const chatName = faker.person.fullName();
  const timeOfLastMessage = faker.date.recent({ days: 5 });
  const lastMessage = faker.lorem.sentence();

  return {
    chatName: chatName,
    timeOfLastMessage: timeOfLastMessage,
    lastMessage: lastMessage,
  };
}

function HomeView() {
  const [currentChatName, setCurrentChatName] = useState("");

  const conversationComponents = useMemo(() => {
    const conversationPreviews: JSX.Element[] = [];
    for (let i = 0; i < 20; i++) {
      const dummy = CreateDummyConversation();
      conversationPreviews.push(
        <ConversationPreview
          key={i}
          conversation={dummy}
          onClick={() => {
            setCurrentChatName(dummy.chatName);
          }}
        />
      );
    }
    conversationPreviews.sort(
      (a, b) =>
        b.props.conversation.timeOfLastMessage.getTime() -
        a.props.conversation.timeOfLastMessage.getTime()
    );
    return conversationPreviews;
  }, []);

  useEffect(() => {
    if (conversationComponents.length > 0) {
      setCurrentChatName(conversationComponents[0].props.conversation.chatName);
    }
  }, [conversationComponents]);

  // setCurrentChatName(conversationComponents[0].props.conversation.chatName);

  // All texts within a conversation
  return (
    <div className="conversations-view-main-container">
      <div className="contacts-container">
        <h2>Messages</h2>
        {conversationComponents}
      </div>
      <div className="conversations-container">
        <h2>{currentChatName}</h2>
        <MessagesView />
      </div>
    </div>
  );
}

export { ConversationPreview as ConversationComponent };
export default HomeView;
