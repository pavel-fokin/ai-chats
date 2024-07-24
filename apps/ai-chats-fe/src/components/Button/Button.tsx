import { forwardRef } from 'react';

import { Button as RadixButton } from '@radix-ui/themes';

interface ButtonElement extends React.ElementRef<typeof RadixButton> {}
interface ButtonProps extends React.ComponentProps<typeof RadixButton> {}

export const Button = forwardRef<ButtonElement, ButtonProps>((props, ref) => {
  return <RadixButton {...props} ref={ref} />;
});
