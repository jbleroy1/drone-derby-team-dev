import React, { createContext, useContext, useEffect, useState } from 'react';
import { signInAnonymously, onAuthStateChanged, User, UserCredential } from 'firebase/auth';

import { auth } from '../firebase/config';

interface AppContextInterface {
  user: User | null;
  signIn: () => Promise<UserCredential>;
}

type AuthProviderProps = {
  children: React.ReactNode;
};

const UserContext = createContext<AppContextInterface>({} as AppContextInterface);

export const AuthContextProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [user, setUser] = useState<User | null>(null);

  const signIn = (): Promise<UserCredential> => {
    return signInAnonymously(auth);
  };

  useEffect(() => {
    const unsubscribe = onAuthStateChanged(auth, (currentUser) => {
      setUser(currentUser);
    });

    return () => {
      unsubscribe();
    };
  }, []);

  return <UserContext.Provider value={{ user, signIn }}>{children}</UserContext.Provider>;
};

export const UserAuth = () => {
  return useContext(UserContext);
};
