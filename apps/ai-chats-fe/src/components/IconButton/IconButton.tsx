import { forwardRef } from 'react';

import { IconButton as RadixIconButton } from '@radix-ui/themes';

interface IconButtonElement extends React.ElementRef<typeof RadixIconButton> {}
interface IconButtonProps
  extends React.ComponentProps<typeof RadixIconButton> {}

export const IconButton = forwardRef<IconButtonElement, IconButtonProps>(
  (props, ref) => {
    return <RadixIconButton {...props} ref={ref} />;
  },
);
