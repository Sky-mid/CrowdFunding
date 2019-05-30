package utils;

import java.io.IOException;
import java.security.NoSuchAlgorithmException;
import java.security.NoSuchProviderException;
import java.security.spec.InvalidKeySpecException;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.TimeoutException;

import org.hyperledger.fabric.sdk.exception.CryptoException;
import org.hyperledger.fabric.sdk.exception.InvalidArgumentException;
import org.hyperledger.fabric.sdk.exception.ProposalException;
import org.hyperledger.fabric.sdk.exception.TransactionException;

import fabric.ChaincodeManager;


public class main {
	public static void Main(String[] args) throws CryptoException, InvalidArgumentException, NoSuchAlgorithmException, NoSuchProviderException, InvalidKeySpecException, TransactionException, IOException, ProposalException, InterruptedException, ExecutionException, TimeoutException{
		ChaincodeManager  fm = FabricManager.obtain().getManager();
		
		String[] test ={"Mike","50","Amy","20"};
		
		fm.invoke("invoke", test);
		
		
	}
}
