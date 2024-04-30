const SendMessage = async (text: string) => {
    const resp = await fetch('/api/messages', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ text }),
      });
    return await resp.json();
}

export { SendMessage };