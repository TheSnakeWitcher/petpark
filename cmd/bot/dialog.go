package main

import "github.com/go-telegram/ui/dialog"

var dialogNodes = []dialog.Node{
		{   
		    ID: "homeNode",
		    Text: "Opciones disponibles:",
		    Keyboard: [][]dialog.Button{
		        {
		            {Text: "Adoptar", NodeID: "adoptNode"},
		            {Text: "Dar en adopcion", NodeID: "giveInAdoptionNode"},
		        },
		        {
		            {Text: "Ir a sitio web", URL: "https://github.com/go-telegram/ui"},
		        },
		    },
		},
	    {
	        ID: " adoptNode",
	        Text: "adoptar",
	        Keyboard: [][]dialog.Button{
	            {
	                {Text: "volver al inicio" , NodeID: "homeNode"},
	                {Text: "dar en adopcoin" , NodeID: "giveInAdoptionNode"},
	            },
	        },
	    },
	    {
	        ID: " giveInAdoptionNode",
	        Text: "dar en adopcion",
	        Keyboard: [][]dialog.Button{
	            {
	                {Text: "volver al inicio" , NodeID: "homeNode"},
	                {Text: "adoptar" , NodeID: "adoptNode"},
	            },
	        },
	    },
	}
