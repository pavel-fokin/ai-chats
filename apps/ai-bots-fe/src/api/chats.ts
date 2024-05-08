import { Chat, Message } from 'types';
import { fetchData, postData } from './base';

type PostChatsResponse = {
    data: {
        id: string;
    };
};

type GetChatsResponse = {
    data: {
        chats?: Chat[];
    }
};

type GetMessagesResponse = {
    data: {
        messages: Message[];
    };
};

type PostMessagesResponse = {
    data: {
        message: Message;
    };
};


const postChats = async (): Promise<PostChatsResponse> => {
    return await postData('/api/chats', {});
}

const fetchChats = async (): Promise<GetChatsResponse> => {
    const payload = await fetchData<GetChatsResponse>('/api/chats');
    return payload;
}

const fetchMessages = async (chatId: string): Promise<GetMessagesResponse> => {
    return await fetchData(`/api/chats/${chatId}/messages`);
}

const postMessages = async (chatId: string, message: Message): Promise<PostMessagesResponse> => {
    return await postData(`/api/chats/${chatId}/messages`, message);
}

export { postChats, fetchChats, fetchMessages, postMessages }