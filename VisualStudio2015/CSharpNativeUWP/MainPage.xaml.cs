using System;
using System.Diagnostics;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Runtime.InteropServices.WindowsRuntime;
using Windows.Storage;
using Windows.Storage.Search;
using Windows.Foundation;
using Windows.Foundation.Collections;
using Windows.UI.Xaml;
using Windows.UI.Xaml.Controls;
using Windows.UI.Xaml.Controls.Primitives;
using Windows.UI.Xaml.Data;
using Windows.UI.Xaml.Input;
using Windows.UI.Xaml.Media;
using Windows.UI.Xaml.Navigation;

using SciMark2;

// The Blank Page item template is documented at http://go.microsoft.com/fwlink/?LinkId=402352&clcid=0x409

namespace CSharpNativeUWP
{
    /// <summary>
    /// An empty page that can be used on its own or navigated to within a Frame.
    /// </summary>
    public sealed partial class MainPage : Page
    {
        public MainPage()
        {
            this.InitializeComponent();
        }

        private async void Button_Click(object sender, RoutedEventArgs e)
        {
            StorageFolder saveFolder = await KnownFolders.GetFolderForUserAsync(null /* current user */, KnownFolderId.DocumentsLibrary);
            var sampleFile =    await saveFolder.CreateFileAsync("UWPResults.txt", Windows.Storage.CreationCollisionOption.OpenIfExists);
            int NumTimes = 50;
            for (int iTime = 0; iTime < NumTimes; iTime++)
            {
                var value = CommandLine.RunTest(new string[0]);
                var msg = string.Format("Win10VM,UWP2015,{0:F2}\n", value);
                Debug.WriteLine(msg.TrimEnd());
                await Windows.Storage.FileIO.AppendTextAsync(sampleFile, msg);
            }
            Debug.WriteLine("finish tests");
            Application.Current.Exit();
        }
    }
}
