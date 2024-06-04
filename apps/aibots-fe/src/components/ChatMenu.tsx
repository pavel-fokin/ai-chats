import { Button, DropdownMenu, Flex } from '@radix-ui/themes';

import {
    Trash as DeleteIcon,
    SlidersHorizontal as SettingsIcon
} from "@phosphor-icons/react";

import { useChat } from "hooks";

type ChatMenuProps = {
    chatId: string;
}

export const ChatMenu = ({ chatId }: ChatMenuProps) => {
    const { data: chat } = useChat(chatId);

    const handleDelete = () => {
        console.log('Delete chat', chatId);
    }

    return (
        <DropdownMenu.Root >
            <DropdownMenu.Trigger>
                <Button
                    variant="ghost"
                    size="3"
                    highContrast
                >
                    <span
                        style={{
                            overflow: "hidden",
                            textOverflow: "ellipsis",
                            whiteSpace: "nowrap",
                            maxWidth: "192px",
                        }}
                    >
                        {chat?.title || 'Chat'}
                    </span>
                    <DropdownMenu.TriggerIcon />
                </Button>
            </DropdownMenu.Trigger>
            <DropdownMenu.Content style={{ minWidth: "128px" }} >
                <DropdownMenu.Item shortcut="">
                    <Flex direction="row" align="center" justify="between" width="100%">
                        Configure <SettingsIcon size="16" />
                    </Flex>
                </DropdownMenu.Item>
                <DropdownMenu.Separator />
                <DropdownMenu.Item color="tomato" onClick={handleDelete}>
                    <Flex direction="row" align="center" justify="between" width="100%">
                        Delete <DeleteIcon size="16" />
                    </Flex>
                </DropdownMenu.Item>
            </DropdownMenu.Content>
        </DropdownMenu.Root>
    );
}