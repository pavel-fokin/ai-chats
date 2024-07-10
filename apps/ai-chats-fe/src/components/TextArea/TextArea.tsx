import styles from './TextArea.module.css';

type TextAreaProps = {
  placeholder: string;
  value: string;
  onChange: (event: React.ChangeEvent<HTMLTextAreaElement>) => void;
};

import React, { useState } from 'react';

export const TextArea: React.FC<TextAreaProps> = (props) => {
  const [rows, setRows] = useState(1);
  // const [text, setText] = useState('');

  const handleChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
    // const { value } = event.target;
    // setText(value);

    // Count the number of line breaks to adjust the rows
    const lineBreaks = (props.value.match(/\n/g) || []).length + 1;
    setRows(lineBreaks + 1);

    props.onChange(event);
  };

  return (
    <textarea
      rows={rows}
      value={props.value}
      onChange={handleChange}
      className={styles.TextArea}
      placeholder={props.placeholder}
    />
  );
};
