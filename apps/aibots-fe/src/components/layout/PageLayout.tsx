import { Root, Aside } from 'components/layout';
import { Navbar } from 'components';


type PageLayoutProps = {
    children: React.ReactNode;
}

export const PageLayout: React.FC<PageLayoutProps>= ({ children }) => {
    return (
        <Root>
            <Aside>
                <Navbar />
            </Aside>
            {children}
        </Root>
    )
}