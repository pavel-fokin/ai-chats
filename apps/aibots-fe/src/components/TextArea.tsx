import TextareaAutosize from 'react-textarea-autosize';

import styles from './TextArea.module.css';

type TextAreaProps = {
  placeholder: string;
  value: string;
  onChange: (event: React.ChangeEvent<HTMLTextAreaElement>) => void;
};

const TextArea = (props: TextAreaProps) => {
  return (
    <TextareaAutosize
      value={props.value}
      onChange={props.onChange}
      className={styles.TextArea}
      placeholder={props.placeholder}
    />
  );
};

export { TextArea };
