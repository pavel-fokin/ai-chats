import { forwardRef } from 'react';

import { TextField } from '@radix-ui/themes'

interface TextInputElement extends React.ElementRef<typeof TextField.Root> {}
interface TextInputProps extends React.ComponentProps<typeof TextField.Root> {}

export const TextInput = forwardRef<TextInputElement, TextInputProps>((props, ref) => {
  return <TextField.Root {...props} ref={ref} />;
});