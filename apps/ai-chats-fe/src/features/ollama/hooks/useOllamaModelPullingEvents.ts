import { useEffect, useRef, useState } from 'react';

import { useInvalidateOllamaModels } from 'hooks';

import { EventTypes, OllamaModel } from 'types';

type EventHandler = (event: MessageEvent) => void;
const eventHandlers = new Map<string, EventHandler>();

interface ProgressEvent {
  status: string;
  total: number;
  completed: number;
}

export const useOllamaModelPullingEvents = (model: OllamaModel) => {
  const invalidateOllamaModels = useInvalidateOllamaModels();
  const [progress, setProgress] = useState<ProgressEvent | null>(null);
  const eventSourceRef = useRef<EventSource | null>(null);
  const accessToken = localStorage.getItem('accessToken') || '';

  eventHandlers.set(EventTypes.OLLAMA_MODEL_PULL_PROGRESS, (event) => {
    const progress: ProgressEvent = JSON.parse(event.data);
    if (progress.status === 'success') {
      invalidateOllamaModels();
    } else if (progress.status === 'error') {
      invalidateOllamaModels();
    } else {
      setProgress(progress);
    }
  });

  useEffect(() => {
    eventSourceRef.current = new EventSource(
      `/api/ollama/models/${model.model}/pulling-events?accessToken=${accessToken}`,
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

  return { progress };
};
