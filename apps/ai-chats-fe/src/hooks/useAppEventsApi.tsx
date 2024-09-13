import { useEffect, useRef } from 'react';

type EventHandler = (event: Event) => void;
const eventHandlers = new Map<string, EventHandler>();

eventHandlers.set('message', (event) => {
  console.log('Message added:', event);
});

export function useAppEvents() {
  const eventSourceRef = useRef<EventSource | null>(null);
  const accessToken = localStorage.getItem('accessToken') || '';

  useEffect(() => {
    const eventSource = new EventSource(`/api/events/app?accessToken=${accessToken}`);
    eventSourceRef.current = eventSource;

    eventSource.onopen = () => {
      console.log('Connection to server opened.');
    };

    for (const [eventType, eventHandler] of eventHandlers) {
      eventSource.addEventListener(eventType, eventHandler);
    }

    eventSource.onerror = (error) => {
      console.error('EventSource failed:', error);
    };

    return () => {
      eventSource.close();
    };
  });
}
