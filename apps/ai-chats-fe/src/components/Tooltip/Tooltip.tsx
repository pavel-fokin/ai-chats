import { forwardRef } from 'react';

import { Tooltip as RadixTooltip } from '@radix-ui/themes';

interface TooltipElement extends React.ElementRef<typeof RadixTooltip> {}
interface TooltipProps extends React.ComponentProps<typeof RadixTooltip> {}

export const Tooltip = forwardRef<TooltipElement, TooltipProps>(
  (props, ref) => {
    return <RadixTooltip {...props} ref={ref} />;
  }
);
