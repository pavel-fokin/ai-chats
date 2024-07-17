import {
  forwardRef,
  useEffect,
  useImperativeHandle,
  useRef,
  useState,
} from 'react';

import styles from './TextArea.module.css';

interface TextAreaProps {
  placeholder: string;
  value: string;
  onChange: (event: React.ChangeEvent<HTMLTextAreaElement>) => void;
}

export const TextArea = forwardRef<HTMLTextAreaElement, TextAreaProps>(
  (props, ref) => {
    const textareaRef = useRef<HTMLTextAreaElement>(null);
    const [value, setValue] = useState(props.value);

    useImperativeHandle(ref, () => textareaRef.current as HTMLTextAreaElement);

    useEffect(() => {
      if (textareaRef && textareaRef.current) {
        textareaRef.current.style.height = '0px';
        const scrollHeight = textareaRef.current.scrollHeight;
        textareaRef.current.style.height = scrollHeight + 'px';
      }
    }, [value]);

    const handleChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
      setValue(event.target.value);
      // textareaRef.current?.value = event.target.value;
      props.onChange(event);
    };

    return (
      <textarea
        autoFocus
        ref={textareaRef}
        value={value}
        onChange={handleChange}
        className={styles.TextArea}
        placeholder={props.placeholder}
      />
    );
  },
);
