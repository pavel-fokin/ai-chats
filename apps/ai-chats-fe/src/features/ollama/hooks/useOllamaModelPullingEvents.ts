import { useEffect, useRef } from 'react';

import { EventTypes, OllamaModel } from 'types';

type EventHandler = (event: MessageEvent) => void;
const eventHandlers = new Map<string, EventHandler>();

export const useOllamaModelPullingEvents = (model: OllamaModel): void => {
  const eventSourceRef = useRef<EventSource | null>(null);
  const accessToken = localStorage.getItem('accessToken') || '';

  eventHandlers.set(EventTypes.OLLAMA_MODEL_PULLING_PROGRESS, (event) => {
    const progress = JSON.parse(event.data);
    console.log(progress);
  });

  useEffect(() => {
    if (!model.isPulling) {
      return;
    }

    eventSourceRef.current = new EventSource(
      `/api/ollama/models/${model.model}/pulling-events?accessToken=${accessToken}`
    );

    for (const [eventType, eventHandler] of eventHandlers) {
      eventSourceRef.current.addEventListener(eventType, eventHandler);
    }

    eventSourceRef.current.onerror = (error) => {
      console.error('EventSource failed:', error);
    };

    return () => {
      eventSourceRef.current?.close();
    };
  }, [accessToken, model]);
};
