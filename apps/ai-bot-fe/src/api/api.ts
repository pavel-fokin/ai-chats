export type Message = {
  who: string;
  text: string;
};


const SendMessage = async (msg: Message) => {
    const resp = await fetch('/api/messages', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ text: msg.text }),
      });
    return await resp.json();
}

export { SendMessage };