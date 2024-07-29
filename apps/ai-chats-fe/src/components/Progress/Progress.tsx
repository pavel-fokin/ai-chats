import { forwardRef } from 'react';

import { Progress as RadixProgress } from '@radix-ui/themes';

interface ProgressElement extends React.ElementRef<typeof RadixProgress> {}
interface ProgressProps extends React.ComponentProps<typeof RadixProgress> {}

export const Progress = forwardRef<ProgressElement, ProgressProps>(
  (props, ref) => {
    return <RadixProgress highContrast {...props} ref={ref} />;
  }
);
